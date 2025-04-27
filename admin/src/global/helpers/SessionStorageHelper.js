export default class SessionsStorageHelper {
    constructor() {
        this.storage = window.sessionStorage;
    }

    setItem(key, value) {
        this.storage.setItem(key, JSON.stringify(value));
    }

    getItem(key) {
        return JSON.parse(this.storage.getItem(key));
    }

    removeItem(key) {
        this.storage.removeItem(key);
    }

    clear() {
        this.storage.clear();
    }
}