from CaptchaParser import CaptchaParser
from PIL import Image

img = Image.open("api/login/captcha_student.bmp")
parser = CaptchaParser()
captcha = parser.getCaptcha(img)

print captcha