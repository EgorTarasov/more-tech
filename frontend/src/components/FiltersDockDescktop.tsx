import { useEffect, useState } from 'react';
import { useStores } from '../hooks/useStores';
import { observer } from 'mobx-react-lite';

import { Card, Col, Input, Row, Segmented, Typography } from 'antd';
import { SearchProps } from 'antd/es/input';
import { AudioOutlined } from '@ant-design/icons';
import { Button } from '@admiral-ds/react-ui';
import { IFilter } from '../models/Filters';
import { Button as AdmiralButton } from '@admiral-ds/react-ui';
import { SegmentedValue } from 'antd/es/segmented';

const { Search } = Input;

const FiltersDockDescktop = observer(() => {
    const { rootStore } = useStores();
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
        <Card className='filters-dock-descktop'>
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
                    <Segmented
                        onChange={(value: SegmentedValue) => {
                            rootStore.setAtmsShown(value === 'Банкоматы');
                        }}
                        size='large'
                        options={['Офисы', 'Банкоматы']}
                    />
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
                    <AdmiralButton
                        onClick={() => {
                            rootStore.setFilters(filters);
                            rootStore.setFiltersDescktopShown(false);
                        }}
                    >
                        Показать результаты
                    </AdmiralButton>

                    <AdmiralButton
                        onClick={() => rootStore.setFiltersDescktopShown(false)}
                        appearance='secondary'
                    >
                        Закрыть фильтры
                    </AdmiralButton>
                </Col>
            </div>
        </Card>
    );
});

export default FiltersDockDescktop;
