import catppuccin from '@catppuccin/daisyui'

module.exports = {
  content: ['./src/routes/**/*.{svelte,js,ts}'],
  plugins: [require('daisyui')],
  daisyui: {
    themes: [
      'light',
      'dark',
      'night',
      catppuccin('macchiato'),
    ],
  },
};
