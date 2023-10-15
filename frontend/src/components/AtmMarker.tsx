import { IAtm } from '../api/models/IAtm';
import atmSvg from '../assets/atm.svg';
import { useStores } from '../hooks/useStores';

type Props = {
    atm: IAtm;
};

const AtmMarker = ({ atm }: Props) => {
    const { rootStore } = useStores();

    return (
        <div
            onClick={() => {
                rootStore.setSelectedAtm(atm);
                rootStore.fetchRoute();
            }}
            className='office-marker'
        >
            <img src={atmSvg} alt={'Банкомат'} />
        </div>
    );
};

export default AtmMarker;
