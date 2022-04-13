const categoryMen = document.querySelector('#category-men');
const categoryWomen = document.querySelector('#category-women');

function selectedCategory() {
  categoryMen.classList.toggle('category--active');
  categoryWomen.classList.toggle('category--active');
}

categoryMen.addEventListener('click', selectedCategory);
categoryWomen.addEventListener('click', selectedCategory);

const searchBar = document.querySelector('#searchBar');
