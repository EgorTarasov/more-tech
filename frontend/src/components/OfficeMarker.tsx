import { IDepartment } from '../api/models';
import bear from '../assets/react.svg';

type Props = {
    department: IDepartment;
};

const OfficeMarker = ({ department }: Props) => {
    return (
        <div className='office-marker'>
            <img src={bear} alt={department.shortName} />
        </div>
    );
};

export default OfficeMarker;
