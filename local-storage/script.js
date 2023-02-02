// localStorage.setItem('theme', 'dark');
// localStorage.setItem('font', 'bold');

// localStorage.removeItem('font')
// localStorage.clear();

// const theme = localStorage.getItem('theme');
// console.log(theme);

// const userSetting = {
//     theme: 'dark',
//     font: 'light',
//     score: 100
// }
// const settingString = JSON.stringify(userSetting);
// localStorage.setItem('userSetting', settingString);
// const uSetting = JSON.parse(localStorage.getItem('userSetting'))

// console.log(uSetting.score);

/** USE CASES
 * cart
 * user settings
 * filters
 * from data
 * score
 */

// Get local storage size
// let size = 0;
// for (let i = 0; i < localStorage.length; i++) {
//     size += localStorage.getItem(localStorage.key(i)).length;
// }
// console.log(size + " bytes");

const themeSelector = document.querySelector("#themeSelector");
const theme = localStorage.getItem("theme");

changeTheme(theme);

themeSelector.addEventListener("change", (e) => {
  localStorage.setItem("theme", e.target.value);
  changeTheme(e.target.value);
});

function changeTheme(theme) {
  if (theme === "dark") document.body.style.backgroundColor = "#222";
  else if (theme === "light") document.body.style.backgroundColor = "#efefef";
  else document.body.style.backgroundColor = "#fff";
}

window.addEventListener("storage", (e) => {
  if (e.key === "theme") {
    changeTheme(e.newValue);
    themeSelector.value = e.newValue;
  }
});

// Error Handling
function storageAvailable(type) {
    let storage;
    try {
        storage = window[type];
        const x = '__storage_test__';
        storage.setItem(x, x);
        storage.removeItem(x);
        return true;
    }
    catch (e) {
        return e instanceof DOMException && (
            // everything except Firefox
            e.code === 22 ||
            // Firefox
            e.code === 1014 ||
            // test name field too, because code might not be present
            // everything except Firefox
            e.name === 'QuotaExceededError' ||
            // Firefox
            e.name === 'NS_ERROR_DOM_QUOTA_REACHED') &&
            // acknowledge QuotaExceededError only if there's something already stored
            (storage && storage.length !== 0);
    }
}

if (storageAvailable('localStorage')) {
    // Yippee! We can use localStorage awesomeness
    console.log('Yippee! Thats working.');
  }
  else {
    // Too bad, no localStorage for us
    console.log('OOPS! Thats not working.');
  }

/** Cookies vs Local Storage vs Session Storage 
 * https://dev.to/iggredible/cookies-vs-local-storage-vs-session-storage-3gp3
 * size: localStorage size 5mb > cookies.
 * localStorage normal mode have not expires date bt cookies have expires date.
 * cookies data every request goto server bt localStorage data not goto server.
 */