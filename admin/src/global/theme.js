import { createTheme } from '@mantine/core';
import '../../core/vars.css';

const theme = createTheme({
    autoContrast: true,
    fontFamily: 'var(--font-family)',
    headings: { fontFamily: 'var(font-family-heading)' },
    cssVariablesResolver: () => ({
        variables: {
            '--mantine-primary-color': 'var(--color-primary)',

            '--mantine-primary-color-filled': 'var(--color-primary)',
            '--mantine-primary-color-filled-hover': 'color-mix(in srgb, var(--color-primary) 10%, black)',
            '--mantine-primary-color-light': 'color-mix(in srgb, var(--color-primary) 90%, white)',

            '--mantine-font-family': 'var(--font-family)',
            '--mantine-font-family-headings': 'var(font-family-heading)',

            '--mantine-radius-default': 'var(border-radius)',
            '--mantine-radius-xs': 'calc(var(border-radius) / 4)',
            '--mantine-radius-sm': 'calc(var(border-radius) / 2)',
            '--mantine-radius-lg': 'calc(var(border-radius) * 2)',
            '--mantine-radius-xl': 'calc(var(border-radius) * 4)',



            '--button-hover:': 'var(--color-primary)',
            
            ' --mantine-primary-color-filled':'var(--color-primary)',
            '--mantine-color-anchor': 'var(--color-primary) !important',


            '--mantine-color-blue-0': `var(--color-primary)`,
            '--mantine-color-blue-1': `var(--color-primary)`,
            '--mantine-color-blue-2': `var(--color-primary)`,
            '--mantine-color-blue-3': `var(--color-primary)`,
            '--mantine-color-blue-4': `var(--color-primary)`,
            '--mantine-color-blue-5': `var(--color-primary)`,
            '--mantine-color-blue-6': `var(--color-primary)`,
            '--mantine-color-blue-7': `var(--color-primary)`,
            '--mantine-color-blue-8': `var(--color-primary)`,
            '--mantine-color-blue-9': `var(--color-primary)`,
            '--mantine-color-blue-filled': `var(--color-primary)`,
            '--mantine-color-blue-filled-hover': `color-mix(in srgb, var(--color-primary) 10%, black)`,
            '--mantine-color-blue-light': `color-mix(in srgb, var(--color-primary) 90%, white)`,
            '--mantine-color-blue-light-hover': `color-mix(in srgb, var(--color-primary) 80%, white)`,
            '--mantine-color-blue-light-color': `var(--color-primary)`,
            '--mantine-color-blue-outline': 'var(--color-primary)',
            '--mantine-color-blue-outline-hover': 'color-mix(in srgb, var(--color-primary) 10%, black)',



            '--mantine-primary-color-0': 'var(--mantine-color-blue-0)',
            '--mantine-primary-color-1': 'var(--mantine-color-blue-1)',
            '--mantine-primary-color-2': 'var(--mantine-color-blue-2)',
            '--mantine-primary-color-3': 'var(--mantine-color-blue-3)',
            '--mantine-primary-color-4':' var(--mantine-color-blue-4)',
            '--mantine-primary-color-5': 'var(--mantine-color-blue-5)',
            '--mantine-primary-color-6': 'var(--mantine-color-blue-6)',
            '--mantine-primary-color-7': 'var(--mantine-color-blue-7)',
            '--mantine-primary-color-8': 'var(--mantine-color-blue-8)',
            '--mantine-primary-color-9': 'var(--mantine-color-blue-9)',
            '--mantine-primary-color-light-hover': 'var(--mantine-color-blue-light-hover)',
            '--mantine-primary-color-light-color': 'var(--mantine-color-blue-light-color)',
        },
    }),
});

export default theme;