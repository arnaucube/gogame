const axios = require('axios');
const sleep = require('sleep');

const url = 'http://127.0.0.1:5000';

let newUser = {
	name: 'user00',
	password: 'user00password',
	email: 'user00@email.com'
};
axios.post(url + '/register', newUser)
  .then(function (res) {
    console.log(res.data);
  })
  .catch(function (error) {
    console.log(error.response.data);
  });

// sleep.sleep(3);

let user = {
	email: 'user00@email.com',
	password: 'user00password'
};
axios.post(url + '/login', user)
  .then(function (res) {
    console.log(res.data);
  })
  .catch(function (error) {
    console.log(error.response.data);
  });
