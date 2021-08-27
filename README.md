<p align="center">
 • 
 <a href="#-about-the-project">About</a> • 
 <a href="#-used-technologies">Used Technologies</a> • 
 <a href="#-how-to-contribute-to-the-project">How to contribute to the project?</a> • 
 <a href="#-author">Author</a>
</p>

<hr />

### 💻 About the project

<p align="center">

The project is complete autenticathion service where user can create an account, login with username
and password, change email or password, furthermore can recover his account. The principal ideia was use
Apache Kafka message service to cross information between the services, some routes in the API write in a topic from Kafka, which is consumed from another micro-serivce (you can check it [Here](https://github.com/aawadallak/Email-dispatcher-service)).

When an user is going to be created, a random verify code is generated and send to a Redis database, which will
be auto-excluded in five minutes in purpose to have a low cache of keys on redis, and user can't utilize that code in another moment. The User information is stored in a PostgreSQL database.

The entire application are inside a docker container, you can run it trhough the command: Docker-compose up,
if you don't have docker on computer, you need download it on the official page before continue. Besides this you 
must have the our entire ecosystem, remember to run our another micro-service (check above) 
and you need too Apache Kafka

The project was moved from another repository, and it will be updated here from now.
</p>
<hr />

### 🛠 Used technologies
To develop the micro service was utilized the following tecs:
- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Redis](https://redis.io/)
- [Kafka](https://kafka.apache.org/)


## 💪 How to contribute to the project

1. Do a **fork** of the project.
2. Create a new branch with your changes: `git checkout -b my-feature`
3. Save your changes and create a commit with a message: `git commit -m "feature: My new feature"`
4. Send your changes: `git push origin my-feature`

## 👨🏻‍🎓 Author

<a href="https://github.com/aawadallak">
 <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/74802742?v=4" width="100px;" alt=""/>
</a>

Send me a message and i'll be happy to help you. 😄

[![Youtube Badge](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white&link=http)](mailto:alexandre.awadallak@gmail.com)

[![Linkedin Badge](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white&link=https://www.linkedin.com/in/alexandre-yasser-awadallak-1900951b0/)](https://www.linkedin.com/in/alexandre-awadallak)

[![Telegram Badge](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white&link=https://t.me/aawadallak)](https://t.me/aawadallak)


## 📝 License

This project is under license [MIT](./LICENSE).
