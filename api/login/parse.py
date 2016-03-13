from CaptchaParser import CaptchaParser
from PIL import Image
import sys
reg = sys.argv[1]
img = Image.open("api/login/"+reg+".bmp")
parser = CaptchaParser()
captcha = parser.getCaptcha(img)

print captcha