from bs4 import BeautifulSoup
from CaptchaParser import CaptchaParser
from PIL import Image
import mechanize
import cookielib
import sys

regno = sys.argv[1]
password = sys.argv[2]
br = mechanize.Browser()
br.set_handle_robots(False)
br.set_handle_equiv(True)
br.set_handle_redirect(True)
br.set_handle_referer(True)
cj = cookielib.MozillaCookieJar('api/login/'+regno+'.txt')
br.set_cookiejar(cj)
response = br.open("https://academics.vit.ac.in/student/stud_login.asp")

br.select_form("stud_login")

soup = BeautifulSoup(response.get_data(),'html.parser')
img = soup.find('img', id='imgCaptcha')

br.retrieve("https://academics.vit.ac.in/student/"+img['src'], "api/login/captcha_student.bmp")

img = Image.open("api/login/captcha_student.bmp")
parser = CaptchaParser()
captcha = parser.getCaptcha(img)

br["regno"] = str(regno)
br["passwd"] = str(password)
br["vrfcd"] = str(captcha)

br.method = "POST"
response = br.submit()
cj.save(ignore_discard=True)