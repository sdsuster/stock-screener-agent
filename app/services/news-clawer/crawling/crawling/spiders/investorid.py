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
        paths = ['market']
        for path in paths:
            yield scrapy.Request(url=f"{urls}/{path}", callback=self.parse)

    def parse(self, response):
        # first_col = response.xpath("//main/div[@class='id-grid mx-auto mt-4']")
        # print(first_col)
        items = response.xpath("/html/body/main/div[@class='id-grid mx-auto mt-4']/div[@class='row']/div[@class='col']").css("div.row")

        first_item = items[0]
        
        for i in range(1, len(items) - 1):
            print(i, items[i].css('h4::text').get())
            # print(item)
            # print('++++++++++++++++++')
            # title = item.css("//div.row h3::text").get()
            # url = item.css("a::attr(href)").get()

            # print(title)
            # print(url)