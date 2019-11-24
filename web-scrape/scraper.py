from selenium import webdriver 
from selenium.webdriver.common.keys import Keys 
from bs4 import BeautifulSoup 
from time import sleep
from parsel import Selector 
from dotenv import load_dotenv
from pathlib import Path 
import re 
import csv
import os 

def getProfileLinks(personName: str):
    global_search_input = driver.find_element_by_xpath('//*[@id="ember31"]/input')
    global_search_input.send_keys(personName)
    sleep(2.5)
    global_search_input.send_keys(Keys.RETURN)
    sleep(2.5)
    driver.get(driver.current_url)
    sleep(2.5)
    html_page = driver.page_source 
    soup = BeautifulSoup(html_page, "html.parser")
    profile_links = soup.find_all('a', attrs={'href': re.compile('/in/')})
    return profile_links 

def getProfileData(personName: str, limit: int):
    profile_links = getProfileLinks(personName)
    profile_count = 0 
    profile_dict = {}
    for link in profile_links: 
        url = "https://www.linkedin.com" + link.get("href")
        driver.get(url)
        driver.get(driver.current_url)
        sleep(2.5)

        sel = Selector(text = driver.page_source)

        college = sel.xpath('//*[@id="ember95"]/text()').extract_first()
        if college:
            college = college.strip()
        # check if college is "universidade federal de santa maria"
        if college.casefold() != "universidade federal de santa maria":
            # check if limit reached 
            if profile_count < limit: 
                profile_count += 1
                # go to next profile
                sleep(3.5) 
                continue 
            else: 
                # finish loop 
                break 
        
        name = sel.xpath('//*[@id = "ember45"]/div[2]/div[2]/div[1]/ul[1]/li[1]/text()').extract_first()
        if name:
            name = name.strip()
            profile_dict["name"] = name 

        job_title = sel.xpath( '//*[@id="ember45"]/div[2]/div[2]/div[1]/h2/text()').extract_first()
        if job_title:
            job_title = job_title.strip()
            profile_dict["job_title"] = job_title 

        company = sel.xpath('//*[@id="ember91"]/text()').extract_first()
        if company:
            company = company.strip()
            profile_dict["company"] = company
        
        location = sel.xpath('//*[@id="ember45"]/div[2]/div[2]/div[1]/ul[2]/li[1]/text()').extract_first()
        if location:
            location = location.strip()
            profile_dict["location"] = location 

        profile_dict["url"] = url 
        break 
    return profile_dict



# environment configurations
env_path = Path('.') / '.env'
load_dotenv(dotenv_path=env_path)
linkedin_user = os.getenv("LINKEDIN_USERNAME")
linkedin_password = os.getenv("LINKEDIN_PASSWORD")

# load driver for chrome version 77 
driver = webdriver.Chrome('chromedriver')

# access linkedin pages 
driver.get('https://linkedin.com/login')

# login 
username = driver.find_element_by_id('username')
username.send_keys(linkedin_user)
sleep(2.5)
password = driver.find_element_by_id('password')
password.send_keys(linkedin_password)
sleep(2.5)
login_button = driver.find_element_by_xpath('//*[@type="submit"]')
login_button.click()
sleep(2.5)

# access linkedin feed page 
driver.get('https://www.linkedin.com/feed')
# search for persons based on csv file 
# load csv 
with open('egressos.csv', 'w') as csv_file: 
    csv_reader = csv.DictReader(csv_file, delimiter=',')
    count = 1 
    for row in csv_reader:
        print("search count:", count) 
        profile_dict = getProfileData(row["NOME_PESSOA"], 3)
        if len(profile_dict) > 0:
            profile_dict["curso"] = row["NOME_CURSO"]
            profile_dict["ano_evasao"] = row["ANO_EVASAO"]
            with open('egressos_encontrados.csv') as csv_writer: 
                fields = ['nome', 'job_title', 'company', 'location', 'url', 'curso', 'ano_evasao']
                writer = csv.DictWriter(csv_file, fieldnames=fields)
                writer.writeheader()
                writer.writerow(profile_dict)
        sleep(1.5)
driver.quit()      
