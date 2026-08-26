import json
import logging
import signal
from typing import Any

from confluent_kafka import Consumer, KafkaException


logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s | %(levelname)s | %(name)s | %(message)s",
)

logger = logging.getLogger("news-consumer")

running = True


def shutdown(signum, frame):
    global running
    logger.info("Shutdown signal received")
    running = False


signal.signal(signal.SIGINT, shutdown)
signal.signal(signal.SIGTERM, shutdown)


def handle_message(message: Any) -> None:
    try:
        payload = json.loads(message.value().decode("utf-8"))

        logger.info(
            "Received message | topic=%s partition=%s offset=%s",
            message.topic(),
            message.partition(),
            message.offset(),
        )

        # TODO: process your message here
        logger.info("Payload: %s", payload)

    except json.JSONDecodeError:
        logger.exception("Invalid JSON message")
        # Decide whether this should be skipped or sent to a DLQ.


def main():
    config = {
        "bootstrap.servers": "kafka:9092",
        "group.id": "news-embedding-worker",
        "auto.offset.reset": "earliest",
        "enable.auto.commit": False,
    }

    consumer = Consumer(config)

    consumer.subscribe(["news-embedding"])

    logger.info("Kafka consumer started")

    try:
        while running:
            message = consumer.poll(1.0)

            if message is None:
                continue

            if message.error():
                raise KafkaException(message.error())

            try:
                handle_message(message)

                # Commit only after successful processing.
                consumer.commit(message=message, asynchronous=False)

            except Exception:
                logger.exception(
                    "Failed to process message "
                    "topic=%s partition=%s offset=%s",
                    message.topic(),
                    message.partition(),
                    message.offset(),
                )

    finally:
        logger.info("Closing Kafka consumer")
        consumer.close()


if __name__ == "__main__":
    main()