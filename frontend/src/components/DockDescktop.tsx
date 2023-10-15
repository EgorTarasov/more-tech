import { List } from 'antd-mobile';
import { useStores } from '../hooks/useStores';
import DepartmentGeneral from './DepartmentGeneral';
import DepartmentDetails from './DepartmentDetails';
import { observer } from 'mobx-react-lite';
import { Card } from 'antd';

const DockDesktop = observer(() => {
    const { rootStore } = useStores();

    return (
        <>
            <Card className='dock-descktop'>
                <List>
                    {rootStore.selectedDepartment ? (
                        <DepartmentDetails department={rootStore.selectedDepartment} />
                    ) : (
                        rootStore.filteredDepartments.map((department, index) => (
                            <List.Item key={index}>
                                <DepartmentGeneral department={department} />
                            </List.Item>
                        ))
                    )}
                </List>
            </Card>
        </>
    );
});
export default DockDesktop;
