import { IDepartment } from '../api/models';
import vtb_green from '../assets/vtb_green.svg';
import vtb_orange from '../assets/vtb_orange.svg';
import vtb_red from '../assets/vtb_red.svg';
import { useStores } from '../hooks/useStores';

type Props = {
    department: IDepartment;
};

const OfficeMarker = ({ department }: Props) => {
    const { rootStore } = useStores();

    let img = vtb_green;

    if (department.rating < 4) {
        img = vtb_orange;
    } else if (department.rating < 3.5) {
        img = vtb_red;
    }

    return (
        <div
            onClick={() => {
                rootStore.setSelectedDepartment(department);
            }}
            className='office-marker'
        >
            <img src={img} alt={department.shortName} />
        </div>
    );
};

export default OfficeMarker;
