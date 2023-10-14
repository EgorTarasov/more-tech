import { createBrowserRouter } from 'react-router-dom';
import Departments from '../pages/Departments';

export const router = createBrowserRouter([
    {
        path: '/map',
        element: <Departments />,
    },

    {
        path: '*',
        element: <Departments />,
    },
]);
