from flask import Flask

app = Flask(__name__)

@app.route("/pyapi", methods=['GET'])
def my_app():
    return "Hey bro this is me python", 200


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)