'''
import logging
logging.basicConfig(filename='example.log',level=logging.INFO)
logging.debug('the message should be print in example log')
logging.info('so should there')
logging.warning('and this too')
'''
import logging
import mylib

def main():
	logging.basicConfig(filename='example',filemode='w',level=logging.INFO)
	logging.info('started')
	mylib.do_something()
	logging.info('finished')

if __name__ == '__main__':
	main()
