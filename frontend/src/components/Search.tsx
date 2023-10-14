import { useStores } from '../hooks/useStores';
import searchSvg from '../assets/vitya-the-bear.svg';

const Search = () => {
    const { rootStore } = useStores();

    return (
        <div onClick={() => rootStore.triggerFilter()} className='search'>
            <img src={searchSvg} alt='search' />
        </div>
    );
};

export default Search;
