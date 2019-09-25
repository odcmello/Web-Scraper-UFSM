import csv
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from time import sleep

# Cria um novo csv
writer = csv.writer(open('profiles.csv', 'w'))

writer.writerow(['Name', 'Job Title', 'Company', 'College', 'Location', 'URL', 'Education'])

# Carrega driver do chrome versão 77
driver = webdriver.Chrome('chromedriver')

driver.get('https://www.linkedin.com/login')


username = driver.find_element_by_id('username')

# Envia email para o formulário
username.send_keys('email')
sleep(1)
password = driver.find_element_by_id('password')

# Envia senha para o formulário
password.send_keys('password')
sleep(1)
login_button = driver.find_element_by_xpath('//*[@type="submit"]')

login_button.click()

# Acessa o google
driver.get('https://www.google.com')

search_query = driver.find_element_by_name('q')
search_query.send_keys('site:linkedin.com/in/ AND "Gabriel Gomes Pereira" AND "Universidade Federal de Santa Maria"')

# Simula o Enter
search_query.send_keys(Keys.RETURN)