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
        if(isDarkMode) {
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
        darkMode(btn)
    } else {
        lightMode(btn)
    }
}

window.addEventListener('DOMContentLoaded', (event) => {
    defaultMode();
});