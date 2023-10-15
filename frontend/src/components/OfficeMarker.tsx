import { IDepartment } from '../api/models';
import vtb1 from '../assets/vtb1.svg';
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
            <img src={vtb1} alt={department.shortName} />
        </div>
    );
};

export default OfficeMarker;
