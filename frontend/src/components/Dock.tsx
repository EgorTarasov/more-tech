import { FloatingPanel, List } from 'antd-mobile';
import { useStores } from '../hooks/useStores';
import DepartmentGeneral from './DepartmentGeneral';
import DepartmentDetails from './DepartmentDetails';
import { observer } from 'mobx-react-lite';
import AtmGeneral from './AtmGeneral';

const anchors = [50, window.innerHeight * 0.4, window.innerHeight * 0.8];

const Dock = observer(() => {
    const { rootStore } = useStores();

    return (
        <>
            <FloatingPanel className='dock' anchors={anchors}>
                <List>
                    {rootStore.selectedDepartment ? (
                        <DepartmentDetails department={rootStore.selectedDepartment} />
                    ) : rootStore.isAtmsShown ? (
                        rootStore.atms.map((atm, index) => (
                            <List.Item
                                onClick={() => {
                                    rootStore.setSelectedAtm(atm);
                                    rootStore.fetchRoute();
                                }}
                                key={index}
                            >
                                <AtmGeneral atm={atm} />
                            </List.Item>
                        ))
                    ) : (
                        rootStore.filteredDepartments.map((department, index) => (
                            <List.Item key={index}>
                                <DepartmentGeneral department={department} />
                            </List.Item>
                        ))
                    )}
                </List>
            </FloatingPanel>
        </>
    );
});
export default Dock;
