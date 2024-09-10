<script>
    import { username } from '$lib/stores/username';
    import { goto } from '$app/navigation';

    let name

    async function handleSubmit() {
        console.log('handleSubmit')
        // set the button to loading state
        let button = document.querySelector('button')
        let input = document.querySelector('input')
        button.innerHTML = 'Loading...'
        input.disabled = true
        button.disabled = true
        name = await username.login(name)
        if (name == false) {
            button.innerHTML = 'Start chatting'
            button.disabled = false
            input.disabled = false
            return
        }
        goto('/')
    }
</script>

<h1>Welcome back to Commnet{#if name != ""}, {name} 👋{:else} 👋{/if}</h1>
<p>Enter your name to login:</p>
<input bind:value={name} placeholder="Your name" />
<button on:click={handleSubmit}>Start chatting</button>
<button on:click={() => goto('/')}>/</button>