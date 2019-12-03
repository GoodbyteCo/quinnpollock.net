const isDarkMode = window.matchMedia("(prefers-color-scheme: dark)").matches;
const isLightMode = window.matchMedia("(prefers-color-scheme: light)").matches;
const isNotSpecified = window.matchMedia("(prefers-color-scheme: no-preference)").matches;
const hasNoSupport = !isDarkMode && !isLightMode && !isNotSpecified;

function darkMode(btn) {
    btn.innerHTML = "dark mode: on";
    document.body.className = "dark-mode";	
}

function lightMode(btn) {
    btn.innerHTML = "dark mode: off";
    document.body.className = "";
}

function defaultMode() {
    if(!(CSS.supports("backdrop-filter: invert(1)") || CSS.supports("-webkit-backdrop-filter: invert(1)"))) {
        console.log("firefox")
        var btn = document.getElementById("dark-mode-btn");
        btn.innerHTML = "" 
    } else {
		if(((localStorage.getItem("dark-mode") === null) && (isDarkMode)) || (localStorage.getItem("dark-mode") == 1)) {
            var btn = document.getElementById("dark-mode-btn");
            darkMode(btn)
        }
    }
}

window.matchMedia("(prefers-color-scheme: dark)").addListener(e => {
    var btn = document.getElementById("dark-mode-btn");
    e.matches && darkMode(btn)
});

window.matchMedia("(prefers-color-scheme: light)").addListener(e => {
    var btn = document.getElementById("dark-mode-btn"); 
    e.matches && lightMode(btn)
});

function changeMode() {
    document.getElementById("nav").classList.toggle("closed");
    var btn = document.getElementById("dark-mode-btn");
    
    if(btn.innerHTML == "dark mode: off") {
        darkMode(btn);
		localStorage.setItem("dark-mode", 1);
    } else {
        lightMode(btn);
		localStorage.setItem("dark-mode", 0);
    }
}

window.addEventListener('DOMContentLoaded', (event) => {
    defaultMode();
});