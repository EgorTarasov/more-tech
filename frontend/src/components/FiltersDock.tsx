import { FloatingPanel, FloatingPanelRef } from 'antd-mobile';

import { useEffect, useRef, useState } from 'react';
import { useStores } from '../hooks/useStores';
import { observer } from 'mobx-react-lite';

import { Col, Input, Row, Segmented, Typography } from 'antd';
import { SearchProps } from 'antd/es/input';
import { AudioOutlined } from '@ant-design/icons';
import { Button } from '@admiral-ds/react-ui';
import { IFilter } from '../models/Filters';
import { Button as AdmiralButton } from '@admiral-ds/react-ui';

const { Search } = Input;

const anchors = [0, window.innerHeight - 20];

const FiltersDock = observer(() => {
    const { rootStore } = useStores();
    const ref = useRef<FloatingPanelRef>(null);
    const [filters, setFilters] = useState<IFilter>(rootStore.filters);

    const officeFilters = [
        {
            title: 'Для физлиц',
            active: filters?.special?.person,
            key: 'person',
            onChange: () => {
                setFilters({
                    ...filters,
                    special: { ...filters?.special, person: !filters?.special.person },
                });
            },
        },
        {
            title: 'Для юрлиц',
            active: filters?.special?.juridical,
            key: 'juridical',
            onChange: () => {
                setFilters({
                    ...filters,
                    special: { ...filters?.special, juridical: !filters?.special.juridical },
                });
            },
        },
        {
            title: 'ВТБ "Привелегия"',
            active: filters?.special?.vipZone,
            key: 'vipZone',
            onChange: () => {
                setFilters({
                    ...filters,
                    special: { ...filters?.special, vipZone: !filters?.special.vipZone },
                });
            },
        },
        {
            title: 'Категория "Прайм"',
            active: filters?.special?.Prime,
            key: 'Prime',
            onChange: () => {
                setFilters({
                    ...filters,
                    special: { ...filters?.special, Prime: !filters?.special.Prime },
                });
            },
        },
        {
            title: 'Для маломобильных',
            active: filters?.special?.ramp,
            key: 'ramp',
            onChange: () => {
                setFilters({
                    ...filters,
                    special: { ...filters?.special, ramp: !filters?.special.ramp },
                });
            },
        },
        {
            title: 'Оценка от 4.0',
            active: filters?.raitingMoreThan4,
            key: 'raitingMoreThan4',
            onChange: () => {
                setFilters({
                    ...filters,
                    raitingMoreThan4: !filters?.raitingMoreThan4,
                });
            },
        },
        {
            title: 'Оценка от 4.5',
            active: filters?.raitingMoreThan45,
            key: 'raitingMoreThan45',
            onChange: () => {
                setFilters({
                    ...filters,
                    raitingMoreThan45: !filters?.raitingMoreThan45,
                });
            },
        },
    ];

    useEffect(() => {
        setFilters(rootStore.filters);
    }, [rootStore.filters]);

    useEffect(() => {
        if (rootStore.openFilterTrigger !== null) {
            ref.current?.setHeight(window.innerHeight - 20);
        }
    }, [rootStore.openFilterTrigger]);

    const onSearch: SearchProps['onSearch'] = (value: string) => {
        rootStore.searchML(value);
    };

    const suffix = (
        <AudioOutlined
            style={{
                fontSize: 16,
                color: '#0062ff',
            }}
        />
    );

    return (
        <FloatingPanel ref={ref} className='filters-dock' anchors={anchors}>
            <div style={{ padding: '0px 12px' }}>
                <Row>
                    <Col span={24}>
                        <Search
                            className='filters-dock__search'
                            placeholder='Расскажите, что вас интересует?'
                            allowClear
                            size='large'
                            onSearch={onSearch}
                            suffix={suffix}
                            loading={rootStore.isSearchLoading}
                        />
                    </Col>
                </Row>

                <Row>
                    <Typography.Title level={4}>Фильтры</Typography.Title>
                </Row>

                <Row>
                    <Segmented size='large' options={['Офисы', 'Банкоматы']} />
                </Row>

                <Row>
                    <div className='filters'>
                        {officeFilters.map((filter, index) => {
                            const active = filter.active;

                            return (
                                <Button
                                    key={index}
                                    className={`filters__button ${
                                        active ? 'filters__button_active' : ''
                                    }`}
                                    dimension='m'
                                    onClick={filter.onChange}
                                >
                                    {filter.title}
                                </Button>
                            );
                        })}
                    </div>
                </Row>
            </div>

            <div className='department__details__actions'>
                <Col>
                    <AdmiralButton style={{ opacity: 0 }}>Записаться в отделение</AdmiralButton>

                    <AdmiralButton
                        onClick={() => {
                            rootStore.setFilters(filters);
                            ref.current?.setHeight(0);
                        }}
                    >
                        Показать результаты
                    </AdmiralButton>
                </Col>
            </div>
        </FloatingPanel>
    );
});

export default FiltersDock;
