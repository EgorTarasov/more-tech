import { FloatingPanel, FloatingPanelRef } from 'antd-mobile';

import { useEffect, useRef } from 'react';
import { useStores } from '../hooks/useStores';
import { observer } from 'mobx-react-lite';

const anchors = [0, window.innerHeight - 20];

const AppointmentDock = observer(() => {
    const { rootStore } = useStores();
    const ref = useRef<FloatingPanelRef>(null);

    useEffect(() => {
        if (rootStore.openAppointmentTrigger !== null) {
            console.log('openAppointmentTrigger', rootStore.openAppointmentTrigger);

            ref.current?.setHeight(window.innerHeight - 20);
        }
    }, [rootStore.openAppointmentTrigger]);

    return (
        <FloatingPanel ref={ref} className='appointment-dock' anchors={anchors}>
            123
        </FloatingPanel>
    );
});

export default AppointmentDock;
