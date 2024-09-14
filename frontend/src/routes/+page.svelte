<script>
    import { Hello } from '$lib/wailsjs/go/main/App.js';
    import { username, verified } from '$lib/stores/username.js';

    let text = 'Press the button to fetch data from the backend';

    async function helloWorld() {
        text = await Hello();
    }
</script>

<button class="btn" on:click={helloWorld}>{ text }</button>

<p class="mt-4">Username: { $username }</p>
<p>Verified: { $verified }</p>


