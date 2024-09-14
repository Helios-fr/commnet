<script>
    import { Hello } from '$lib/wailsjs/go/main/App.js';
    import { username, verified } from '$lib/stores/username.js';
    import { onMount } from 'svelte';

    onMount(async () => {
        console.log($username);
        if ($username === '') {
            await username.setDefault();
        }
    });

    let text = 'Press the button to fetch data from the backend';

    async function helloWorld() {
        text = await Hello();
    }
</script>

<button class="btn" on:click={helloWorld}>{ text }</button>

<p class="mt-4">Username: { $username }</p>
<p>Verified: { $verified }</p>

<h2 class="mt-8 mb-2 text-2xl font-bold">Buttons</h2>

<div class="grid grid-cols-2 gap-2 md:grid-cols-4 lg:grid-cols-8">
    <button class="btn">Default</button>
    <button class="btn btn-primary">Primary</button>
    <button class="btn btn-secondary">Secondary</button>
    <button class="btn btn-accent">Accent</button>
    <button class="btn btn-info">Info</button>
    <button class="btn btn-success">Success</button>
    <button class="btn btn-warning">Warning</button>
    <button class="btn btn-error">Error</button>
</div>
