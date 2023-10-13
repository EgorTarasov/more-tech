import React, { useEffect, useRef, useState } from 'react';
import ReactDOM from 'react-dom';

const Departments = () => {
    const [YMaps, setYMaps] = useState(<div />);
    const map = useRef(null);

    useEffect(() => {
        (async () => {
            try {
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
                    YMapMarker,
                    YMapControls,
                } = reactify.module(ymaps3);
                const { YMapZoomControl, YMapGeolocationControl } = reactify.module(
                    await ymaps3.import('@yandex/ymaps3-controls@0.0.1')
                );

                setYMaps(() => (
                    <YMap
                        location={{
                            center: [37.623082, 55.75254],
                            zoom: 9,
                        }}
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

                        <YMapMarker coordinates={[37.623082, 55.75254]} draggable={true}>
                            <section>
                                <p>Этот заголовок можно перетаскивать</p>
                            </section>
                        </YMapMarker>
                    </YMap>
                ));
            } catch (e) {
                console.log(e);

                setYMaps(<div />);
            }
        })();
    }, []);

    return <div style={{ width: '100%', height: '100vh' }}>{YMaps}</div>;
};

export default Departments;
