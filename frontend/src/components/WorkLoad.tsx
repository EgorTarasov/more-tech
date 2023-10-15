import { Button, Row } from 'antd';
import { IWorkload } from '../api/models';
import { useState } from 'react';
import { BarChart, ResponsiveContainer, Bar, XAxis } from 'recharts';

type Props = {
    workLoad: IWorkload[] | undefined;
};

const WorkLoad = ({ workLoad }: Props) => {
    const [activeDay, setActiveDay] = useState<number>(0);

    console.log(workLoad);

    return (
        <>
            <Row>
                {workLoad && (
                    <div style={{ width: '100%', height: '200px' }}>
                        <ResponsiveContainer width='100%' height='100%'>
                            <BarChart
                                width={150}
                                height={40}
                                data={workLoad[activeDay].loadHours?.map((schedule) => {
                                    return {
                                        name: schedule.hour.split('-')[0],
                                        uv: schedule.load,
                                    };
                                })}
                            >
                                <XAxis dataKey='name' />
                                <Bar dataKey='uv' fill='#0062ff' />
                            </BarChart>
                        </ResponsiveContainer>
                    </div>
                )}
            </Row>
            <Row justify={'space-between'}>
                {workLoad?.map((day, index) => (
                    <Button
                        key={index}
                        onClick={() => setActiveDay(index)}
                        style={{
                            backgroundColor: activeDay === index ? '#EBEDF5' : '#fff',
                            color: activeDay === index ? '#6b7683' : '#000',
                            border: 'none',
                            padding: 10,
                            marginRight: 5,
                        }}
                    >
                        {day.day}
                    </Button>
                ))}
            </Row>
        </>
    );
};

export default WorkLoad;
