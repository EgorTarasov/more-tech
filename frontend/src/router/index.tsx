import { createBrowserRouter } from 'react-router-dom';
import Map from '../pages/Map';

export const router = createBrowserRouter([
    {
        path: '/map',
        element: <Map />,
    },

    {
        path: '*',
        element: <Map />,
    },
]);
