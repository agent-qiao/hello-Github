#使用flash包输出网址

from flask import Flask, render_template,request


def search4letters(phrase:str, letters:str ='aeiou')-> set:
	return set(letters).intersection(set(phrase))

app = Flask(__name__)

@app.route('/')
def hello()->str:
	return 'Hello world form Flask'

@app.route('/entry')
def entry_page()->'html':
	return render_template('entry.html',the_title ='welcome to search4letters on the web')

def log_request(req:'flask_request',res:str)->none:
	with open('vsearch.log','a') as log:
		print(req,res,file=log)

@app.route('/search4',methods=['POST'])
def do_search()->str:
	phrase = request.form['phrase']
	letters= request.form['letters']
	title  = 'here are your result'
	result = str(search4letters(phrase,letters))
	return render_template('results.html',the_title=title,
		the_phrase=phrase,the_letters=letters,the_results=result)

@app.route('/viewlog')
def view_the_log()->str:
	with open(vsearch.log) as log:
		contents = log.read()
	return contents

if __name__ == '__main__':
	app.run(debug = True) 

