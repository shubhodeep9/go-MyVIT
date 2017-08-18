import sys
from PIL import Image
from base64 import decodestring
from parser import CaptchaParse

imagestr = sys.argv[1]
with open("foo.png","wb") as f:
    f.write(decodestring(imagestr))

f.close()
img = Image.open("foo.png")
captcha = CaptchaParse(img)

with open("output.txt","wb") as f2:
    f2.write(captcha)

f2.close()

