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

function onDocumentReady(callback) {
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', callback);
    } else {
        callback();
    }
}

function onHTMXLoad(callback) {
    document.body.addEventListener('htmx:load', callback);
}

function onHTMXError(callback) {
    document.body.addEventListener('htmx:error', callback);
}

function notify(text) {
    UIkit.notification(text, { pos: 'bottom-right' });
}

function notifySuccess(text) {
    UIkit.notification(text, { pos: 'bottom-right', status: 'success' });
}

function notifyInfo(text) {
    UIkit.notification(text, { pos: 'bottom-right', status: 'primary' });
}

function notifyError(text) {
    UIkit.notification(text, { pos: 'bottom-right', status: 'danger' });
}

function errorHandler(error, context = {}, skipNotify = false) {
    console.error({ error, context });
    const message = error instanceof Error ? error.message : error;
    !skipNotify && notifyError(message);
}