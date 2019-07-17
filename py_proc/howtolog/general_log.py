import logging

class GeneralLog:

    def __init__(self,recorder):
        # create the logger which is interface for our calling the class GeneralLog
        self.logger         = logging.getLogger(recorder)
        # set logger level and format
        self.level          = logging.DUBUG
        self.fmt            = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')

        # create the log of StreamHandler which will dispaly on commandline
        self.sh             = logging.StreamHandler()
        # set the format of StreamHandler
        self.sh.setFormatter(self.fmt)
        # and the StreamHandler to logger
        self.logger.andHandler(sh)

