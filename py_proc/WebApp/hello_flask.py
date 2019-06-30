#使用flash包输出网址

from flask import Flask, render_template,request,escape


def search4letters(phrase:str, letters:str ='aeiou')-> set:
	return set(letters).intersection(set(phrase))

app = Flask(__name__)

@app.route('/')
def hello()->str:
	return 'Hello world form Flask'

@app.route('/entry')
def entry_page()->'html':
	return render_template('entry.html',the_title ='welcome to search4letters on the web')

def log_request(req:'flask_request',res:str)->None:
	with open('vsearch.log','a') as log:
		print(req.form,res,file=log,end='|')
		print(req.remote_addr,file=log,end='|')
		print(req.user_agent,file=log,end='|')
		print(res,file=log)

@app.route('/search4',methods=['POST'])
def do_search()->str:
	phrase = request.form['phrase']
	letters= request.form['letters']
	title  = 'here are your result'
	results= str(search4letters(phrase,letters))
	log_request(request,results)
	return render_template('results.html',the_title=title,
		the_phrase=phrase,the_letters=letters,the_results=results)

@app.route('/viewlog')
def view_the_log() -> 'html':
	contents = []
	with open('vsearch.log') as log:
		for line in log:
			contents.append([])
			for item in line.split('|'):
				contents[-1].append(escape(item))
	titles = ('Form Date','Remote_addr','User_agent','Results')
	return render_template('viewlog.html',
							the_title = 'view log',
							the_row_titles = titles,
							the_data = contents)

if __name__ == '__main__':
	app.run(debug = True) 

