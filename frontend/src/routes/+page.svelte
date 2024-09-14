<script>
    import { Hello } from '$lib/wailsjs/go/main/App.js';

    let text = 'Press the button to fetch data from the backend';

    async function helloWorld() {
        text = await Hello();
    }
</script>

<button class="btn" on:click={helloWorld}>{ text }</button>
