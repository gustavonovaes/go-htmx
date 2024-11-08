function cleanErrors() {
    document.getElementById('error').innerHTML = '';
}

function getURL() {
    return new URL(window.location.href);
}

function getQueryParam(name, defaultValue = null) {
    const value = getURL().searchParams.get(name);
    return value ? value : defaultValue;
}

function setQueryParam(name, value = true) {
    const url = getURL();
    url.searchParams.set(name, value);
    window.history.pushState({}, '', url.href);
}

function replaceQueryParam(obj = {}) {
    const url = getURL();
    for (const [name, value] of Object.entries(obj)) {
        url.searchParams.set(name, value);
    }
    window.history.replaceState({}, '', url.href);
}

function getStorage(key, defaultValue = null) {
    const value = localStorage.getItem(key);
    return value ? value : defaultValue;
}

function setStorage(key, value = true) {
    localStorage.setItem(key, value);
}

function replaceStorage(obj = {}) {
    for (const [key, value] of Object.entries(obj)) {
        setStorage(key, value);
    }
}