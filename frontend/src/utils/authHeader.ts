import { RegisterResponse } from '../api/models';

export default function authHeader() {
    if (localStorage.getItem('user') == null) {
        return {};
    }
    const user: RegisterResponse = JSON.parse(localStorage.getItem('user') as string);

    if (user && user.bearer_token) {
        return { Authorization: 'Bearer ' + user.bearer_token };
    }

    return {};
}
