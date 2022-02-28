const express = require('express');
const path = require('path');

const app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(express.static(path.join(__dirname, 'public')));

app.get('/home', (req, res) => {
  res.render('landing');
});

app.get('*', (req, res) => {
  res.redirect('/home');
});

app.listen(3000, () => {
  console.log('Serving on port 3000!');
});

module.exports = app;
