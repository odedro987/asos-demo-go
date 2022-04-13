const arrowIcons = document.querySelectorAll('[id^="arrow-icon"]');
const dropdownButtons = document.querySelectorAll('[id^="dropdown-button"]');

dropdownButtons.forEach((button, i) => {
  button.addEventListener('click', () => {
    arrowIcons[i].classList.toggle('arrow-icon');
  });

  button.addEventListener('focusout', () => {
    if (arrowIcons[i].classList) {
      arrowIcons[i].classList.remove('arrow-icon');
    }
  });
});
