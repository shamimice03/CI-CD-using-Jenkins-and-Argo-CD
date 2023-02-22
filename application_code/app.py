from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Start of a new day. Have a nice day!'
