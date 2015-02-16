FROM debian:7.7
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update
RUN echo "postfix postfix/mailname string gmail.com" | debconf-set-selections
RUN echo "postfix postfix/main_mailer_type string 'Internet Site'" | debconf-set-selections
RUN apt-get install -q -y postfix 
ADD run-smtp.sh /run-smtp.sh
RUN chmod +x /run-smtp.sh
CMD /run-smtp.sh
