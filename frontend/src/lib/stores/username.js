import { writable } from 'svelte/store';
import { Verify, DefaultUsername } from '$lib/wailsjs/go/main/App.js';

function createUsername() {
    const { subscribe, set } = writable('');

    async function verify(user) {
        if (user == await DefaultUsername()) {
            verified.set(true);
            set(user);
            return;
        
        }
        console.log ("Verifying username: " + user);
        let resp = await Verify(user);
        if (resp === 'true') {
            verified.set(true);
            set(user);
        } else {
            verified.set(false);
        }
    }

    async function setDefault() {
        let user = await DefaultUsername();
        set(user);
    }

    return {
        subscribe,
        set: (user) => set(user),
        verify: verify,
        setDefault: setDefault
    };
}

export const username = createUsername();
export const verified = writable(false);