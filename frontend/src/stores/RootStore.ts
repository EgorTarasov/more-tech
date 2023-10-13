import { makeAutoObservable, observable } from 'mobx';

export class RootStore {
    public trigger: boolean = false;

    constructor() {
        makeAutoObservable(this, {
            trigger: observable,
        });
    }
}
