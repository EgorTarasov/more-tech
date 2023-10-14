import React from 'react';
import ReactDOM from 'react-dom/client';
import { RouterProvider } from 'react-router-dom';
import { router } from './router';

import './index.scss';
import { ConfigProvider } from 'antd';
import { ThemeProvider } from 'styled-components';
import { DropdownProvider, FontsVTBGroup, LIGHT_THEME } from '@admiral-ds/react-ui';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <ConfigProvider
            theme={{
                token: {},
            }}
        >
            <ThemeProvider theme={LIGHT_THEME}>
                <DropdownProvider>
                    <FontsVTBGroup />
                    <RouterProvider router={router} />
                </DropdownProvider>
            </ThemeProvider>
        </ConfigProvider>
    </React.StrictMode>
);
