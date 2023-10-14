import { IDepartment } from '../api/models';
import vtb from '../assets/vtb.png';

type Props = {
    department: IDepartment;
};

const OfficeMarker = ({ department }: Props) => {
    return (
        <div className='office-marker'>
            <img src={vtb} alt={department.shortName} />
        </div>
    );
};

export default OfficeMarker;
