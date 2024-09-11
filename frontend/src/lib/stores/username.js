import { writable } from "svelte/store";
import { Login, Register, DefaultUsername } from '$lib/wailsjs/go/main/App.js';

function createUsername() {
    const { subscribe, set } = writable("localhost");

    function login(username) {
        Login(username).then(user => {
            if (user.startsWith("Error")) {
                alert(user);
                return false;
            }
            else {
                set(user);
                return user;
            }
        });
    }

    function register(username) {
        Register(username).then(user => {
            if (user.startsWith("Error")) {
                alert(user);
                return false;
            }
            else {
                set(user);
                return user;
            }
        });
    }

    function logout() {
        DefaultUsername().then(user => {
            set(user);
        });
    }

    function getDefaultUsername() {
        DefaultUsername().then(user => {
            return user;
        });
    }

    logout();

    return {
        subscribe,
        set,
        reset: () => set("default"),
        login,
        register,
        logout
    };
}

export const username = createUsername();
export var isLoggedIn = writable(false);