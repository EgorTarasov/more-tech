import { useStores } from '../hooks/useStores';
import searchSvg from '../assets/vitya-the-bear.svg';
import RoadTime from './RoadTime';

const Search = () => {
    const { rootStore } = useStores();

    return (
        <>
            <div onClick={() => rootStore.triggerFilter()} className='search'>
                <img src={searchSvg} alt='search' />
            </div>
            <RoadTime />
        </>
    );
};

export default Search;
