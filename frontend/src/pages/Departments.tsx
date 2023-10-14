import React, { useEffect, useRef, useState } from 'react';
import ReactDOM from 'react-dom';
import { useStores } from '../hooks/useStores';
import { observer } from 'mobx-react-lite';
import OfficeMarker from '../components/OfficeMarker';
import Search from '../components/Search';
import Dock from '../components/Dock';

const Departments = observer(() => {
    const [YMaps, setYMaps] = useState(<div />);
    const map = useRef(null);
    const { rootStore } = useStores();

    useEffect(() => {
        rootStore.fetchDepartments();
    }, [rootStore]);

    useEffect(() => {
        (async () => {
            try {
                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                // @ts-ignore
                const ymaps3 = window.ymaps3;
                const [ymaps3React] = await Promise.all([
                    ymaps3.import('@yandex/ymaps3-reactify'),
                    ymaps3.ready,
                ]);

                const reactify = ymaps3React.reactify.bindTo(React, ReactDOM);

                const {
                    YMap,
                    YMapDefaultSchemeLayer,
                    YMapDefaultFeaturesLayer,
                    YMapControls,
                    YMapMarker,
                    YMapFeature,
                } = reactify.module(ymaps3);
                const { YMapZoomControl, YMapGeolocationControl } = reactify.module(
                    await ymaps3.import('@yandex/ymaps3-controls@0.0.1')
                );
                // const { YMapDefaultMarker } = reactify.module(
                //     await ymaps3.import('@yandex/ymaps3-markers@0.0.1')
                // );

                setYMaps(() => (
                    <YMap
                        location={rootStore.mapLocation}
                        camera={{ tilt: 0, azimuth: 0, duration: 0 }}
                        ref={map}
                    >
                        <YMapDefaultSchemeLayer />
                        <YMapDefaultFeaturesLayer />
                        <YMapControls position='right'>
                            <YMapZoomControl />
                        </YMapControls>
                        <YMapControls position='left'>
                            <YMapGeolocationControl />
                        </YMapControls>
                        {rootStore.departments.map((department) => {
                            console.log(department.location);

                            return (
                                <YMapMarker
                                    key={department._id}
                                    coordinates={[
                                        department.location.coordinates.longitude,
                                        department.location.coordinates?.latitude,
                                    ]}
                                    draggable={false}
                                    position={'center'}
                                >
                                    <OfficeMarker department={department} />
                                </YMapMarker>
                            );
                        })}
                        <YMapFeature {...rootStore.polylyne} />
                    </YMap>
                ));
            } catch (e) {
                console.log(e);

                setYMaps(<div />);
            }
        })();
    }, [rootStore.mapLocation, rootStore.departments, rootStore.polylyne]);

    return (
        <>
            <Search />
            <div style={{ width: '100%', height: '100vh' }}>{YMaps}</div>
            <Dock />
        </>
    );
});

export default Departments;
