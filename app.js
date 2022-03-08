const express = require('express');
const path = require('path');
const ejsMate = require('ejs-mate');

const app = express();

// view engine setup
app.engine('ejs', ejsMate);
app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(express.static(path.join(__dirname, 'public')));

app.get('/home', (req, res) => {
  res.render('landing');
});

app.get('/products', (req, res) => {
  res.render('products');
});

app.get('*', (req, res) => {
  res.redirect('/home');
});

app.listen(3000, () => {
  console.log('Serving on port 3000!');
});

module.exports = app;
