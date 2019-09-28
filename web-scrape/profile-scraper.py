import csv
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from time import sleep
from bs4 import BeautifulSoup
import re

# Cria um novo csv
writer = csv.writer(open('profiles.csv', 'w'))

writer.writerow(['Name', 'Job Title', 'Company', 'College', 'Location', 'URL', 'Education'])

# Carrega driver do chrome versão 77
driver = webdriver.Chrome('chromedriver')

driver.get('https://www.linkedin.com/login')


username = driver.find_element_by_id('username')

# Envia email para o formulário
username.send_keys('ocruzmello@gmail.com')
sleep(1)
password = driver.find_element_by_id('password')

# Envia senha para o formulário
password.send_keys('sistufsm2019')
sleep(1)
login_button = driver.find_element_by_xpath('//*[@type="submit"]')

login_button.click()

driver.get("https://www.linkedin.com/feed")
sleep(3)

global_search_input = driver.find_element_by_xpath('//*[@id="ember34"]/input')
global_search_input.send_keys('Gabriel Gomes Pereira')
sleep(1)

global_search_input.send_keys(Keys.RETURN)
sleep(1.5)

driver.get(driver.current_url)
sleep(1.6)

html_page = driver.page_source

soup = BeautifulSoup(html_page, "html.parser")

for link in soup.find_all('a', attrs={'href': re.compile('/in/')}):
    print(link.get("href"))


sleep(20)
driver.quit()
