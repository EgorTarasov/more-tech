import { IDepartment } from '../api/models';
import bear from '../assets/react.svg';
import { useStores } from '../hooks/useStores';

type Props = {
    department: IDepartment;
};

const OfficeMarker = ({ department }: Props) => {
    const { rootStore } = useStores();

    return (
        <div
            onClick={() => {
                rootStore.setSelectedDepartment(department);
            }}
            className='office-marker'
        >
            <img src={bear} alt={department.shortName} />
        </div>
    );
};

export default OfficeMarker;
