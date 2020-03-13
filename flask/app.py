from flask import Flask, render_template
from flask_cors import CORS
from .views.user import user_router
from api.config import Config
from api.database import db

app = Flask(__name__)
CORS(app)
app.config.from_object('Config')
db.init_app(app)

@app.route("/")
def index():
	return render_template("index.html")

@app.route("result")
def result():


app.run()
