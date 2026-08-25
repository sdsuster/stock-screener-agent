import scrapy
import os

class InvestorIDSpider(scrapy.Spider):
    name = 'investor-id'

    custom_settings = {
        "USER_AGENT": (
            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
            "AppleWebKit/537.36 (KHTML, like Gecko) "
            "Chrome/139.0.0.0 Safari/537.36"
        ),
    }

    async def start(self):
        urls = 'https://investor.id'
        paths = ['corporate-action']
        for path in paths:
            yield scrapy.Request(url=f"{urls}/{path}", callback=self.parse)

    def parse(self, response):
        first_col = response.css(".id-grid>.row>.col")[0]
        items = first_col.xpath("./*")
        for item in items:
            print(item)
            print('++++++++++++++++++')
            title = item.css("div.row h3::text").get()
            url = item.css("a::attr(href)").get()

            # print(title)
            # print(url)