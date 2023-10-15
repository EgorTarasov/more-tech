import { Accordion, AccordionItem } from '@admiral-ds/react-ui';
import { IDepartment } from '../api/models';
import DepartmentGeneral from './DepartmentGeneral';
import { Button, Col, Row, Typography, notification } from 'antd';
import { useStores } from '../hooks/useStores';
import { LeftOutlined } from '@ant-design/icons';
import { Button as AdmiralButton } from '@admiral-ds/react-ui';
import { useEffect, useState } from 'react';
import { observer } from 'mobx-react-lite';
import WorkLoad from './WorkLoad';

type Props = {
    department: IDepartment;
};

const DepartmentDetails = observer(({ department }: Props) => {
    const { rootStore } = useStores();
    const [selectedTimeIndex, setSelectedTimeIndex] = useState<number>(0);
    const [selectedTime, setSelectedTime] = useState<string>('10:00-11:00');
    const [api, contextHolder] = notification.useNotification();
    const [isLoading, setIsLoading] = useState<boolean>(false);

    useEffect(() => {
        rootStore.fetchRoute();
        rootStore.fetchDepartmentDetails();
    }, [rootStore, rootStore.selectedDepartment]);

    return (
        <>
            {contextHolder}
            <div className='department__details'>
                <Row>
                    <Button
                        style={{ marginTop: 10, padding: 0 }}
                        onClick={() => rootStore.setSelectedDepartment(null)}
                        type='text'
                    >
                        <LeftOutlined />
                        Все отделения
                    </Button>
                </Row>

                <DepartmentGeneral department={department} />

                <Accordion>
                    <AccordionItem
                        id='work-time'
                        className='department__details__item'
                        title={
                            <>
                                <svg
                                    width='24'
                                    height='24'
                                    viewBox='0 0 24 24'
                                    fill='none'
                                    xmlns='http://www.w3.org/2000/svg'
                                >
                                    <g clipPath='url(#clip0_170_4491)'>
                                        <path
                                            d='M12 0C5.37188 0 0 5.37188 0 12C0 18.6281 5.37188 24 12 24C18.6281 24 24 18.6281 24 12C24 5.37188 18.6281 0 12 0ZM12 21.9984C6.47812 21.9984 2.00156 17.5219 2.00156 12C2.00156 6.47812 6.47812 2.00156 12 2.00156C17.5219 2.00156 21.9984 6.47812 21.9984 12C21.9984 17.5219 17.5219 21.9984 12 21.9984ZM12.9984 3.99844H10.9969V12L15.4969 16.5L16.9969 15L12.9984 11.0016V3.99844Z'
                                            fill='#1E4BD2'
                                        />
                                    </g>
                                    <defs>
                                        <clipPath id='clip0_170_4491'>
                                            <rect width='24' height='24' fill='white' />
                                        </clipPath>
                                    </defs>
                                </svg>

                                <Typography.Title
                                    style={{
                                        fontSize: 14,
                                        fontWeight: 'bold',
                                        margin: 0,
                                        marginLeft: 10,
                                    }}
                                    level={5}
                                >
                                    Часы работы
                                </Typography.Title>
                            </>
                        }
                    >
                        <Row>
                            <Typography.Text className='base-text'>
                                <b>Для физ. лиц:</b> {department.schedulefl}
                            </Typography.Text>
                        </Row>

                        <Row>
                            <Typography.Text className='base-text'>
                                <b>Для юр. лиц:</b> {department.schedulejurl}
                            </Typography.Text>
                        </Row>
                    </AccordionItem>

                    <AccordionItem
                        id='work-time'
                        className='department__details__item'
                        title={
                            <>
                                <svg
                                    width='24'
                                    height='24'
                                    viewBox='0 0 24 24'
                                    fill='none'
                                    xmlns='http://www.w3.org/2000/svg'
                                >
                                    <path
                                        d='M7.125 12C8.98896 12 10.5 10.489 10.5 8.625C10.5 6.76104 8.98896 5.25 7.125 5.25C5.26104 5.25 3.75 6.76104 3.75 8.625C3.75 10.489 5.26104 12 7.125 12Z'
                                        fill='#1E4BD2'
                                    />
                                    <path
                                        d='M10.9688 13.875C9.64875 13.2047 8.19188 12.9375 7.125 12.9375C5.03531 12.9375 0.75 14.2191 0.75 16.7812V18.75H7.78125V17.9967C7.78125 17.1061 8.15625 16.2131 8.8125 15.4688C9.33609 14.8744 10.0692 14.3227 10.9688 13.875Z'
                                        fill='#1E4BD2'
                                    />
                                    <path
                                        d='M15.9375 13.5C13.4967 13.5 8.625 15.0075 8.625 18V20.25H23.25V18C23.25 15.0075 18.3783 13.5 15.9375 13.5Z'
                                        fill='#1E4BD2'
                                    />
                                    <path
                                        d='M15.9375 12C18.2157 12 20.0625 10.1532 20.0625 7.875C20.0625 5.59683 18.2157 3.75 15.9375 3.75C13.6593 3.75 11.8125 5.59683 11.8125 7.875C11.8125 10.1532 13.6593 12 15.9375 12Z'
                                        fill='#1E4BD2'
                                    />
                                </svg>

                                <Typography.Title
                                    style={{
                                        fontSize: 14,
                                        fontWeight: 'bold',
                                        margin: 0,
                                        marginLeft: 10,
                                    }}
                                    level={5}
                                >
                                    Загруженность
                                </Typography.Title>
                            </>
                        }
                    >
                        <WorkLoad workLoad={rootStore.selectedDepartmentDetails?.workload} />
                    </AccordionItem>

                    <AccordionItem
                        id='work-time'
                        className='department__details__item'
                        title={
                            <>
                                <svg
                                    width='24'
                                    height='24'
                                    viewBox='0 0 24 24'
                                    fill='none'
                                    xmlns='http://www.w3.org/2000/svg'
                                >
                                    <path
                                        d='M10.4021 4.69575C11.2846 4.69575 11.9999 3.98036 11.9999 3.09787C11.9999 2.21539 11.2846 1.5 10.4021 1.5C9.51959 1.5 8.8042 2.21539 8.8042 3.09787C8.8042 3.98036 9.51959 4.69575 10.4021 4.69575Z'
                                        fill='#1E4BD2'
                                    />
                                    <path
                                        d='M21.5682 16.897L19.3209 17.4588L18.422 13.2639C18.3671 13.0127 18.2281 12.7877 18.0282 12.626C17.8282 12.4644 17.5791 12.3757 17.322 12.3746H13.1639L12.861 9.74962H18.0001V8.24962H12.688L12.4942 6.57028C12.4772 6.42351 12.4316 6.28152 12.3598 6.1524C12.288 6.02329 12.1914 5.90959 12.0757 5.81778C11.9599 5.72598 11.8272 5.65787 11.6851 5.61736C11.5431 5.57684 11.3944 5.5647 11.2476 5.58163L9.38501 5.79656L10.3173 13.8746H17.019L18.1796 19.2904L21.9322 18.3522L21.5682 16.897Z'
                                        fill='#1E4BD2'
                                    />
                                    <path
                                        d='M10.5 21.0005C9.19585 21.0003 7.93228 20.5469 6.92547 19.718C5.91866 18.889 5.23118 17.736 4.98063 16.4561C4.73008 15.1763 4.93203 13.8491 5.55194 12.7017C6.17185 11.5543 7.17119 10.658 8.379 10.166L8.20186 8.63086C5.39733 9.58913 3.375 12.2505 3.375 15.3755C3.375 19.3042 6.57127 22.5005 10.5 22.5005C11.7261 22.4999 12.9313 22.183 13.9992 21.5804C15.067 20.9779 15.9614 20.1101 16.5958 19.0609L16.125 16.8755C15.375 19.1255 13.2469 21.0005 10.5 21.0005Z'
                                        fill='#1E4BD2'
                                    />
                                </svg>

                                <Typography.Title
                                    style={{
                                        fontSize: 14,
                                        fontWeight: 'bold',
                                        margin: 0,
                                        marginLeft: 10,
                                    }}
                                    level={5}
                                >
                                    Доступная среда
                                </Typography.Title>
                            </>
                        }
                    >
                        {department.special.ramp ? 'Есть пандус' : 'Нет пандуса'}
                    </AccordionItem>

                    <AccordionItem
                        id='work-time'
                        className='department__details__item'
                        title={
                            <>
                                <svg
                                    width='24'
                                    height='24'
                                    viewBox='0 0 24 24'
                                    fill='none'
                                    xmlns='http://www.w3.org/2000/svg'
                                >
                                    <path
                                        d='M18.375 0.75H5.625C5.32663 0.75 5.04048 0.868526 4.8295 1.0795C4.61853 1.29048 4.5 1.57663 4.5 1.875V17.625C4.5 17.9234 4.61853 18.2095 4.8295 18.4205C5.04048 18.6315 5.32663 18.75 5.625 18.75H18.375C18.6734 18.75 18.9595 18.6315 19.1705 18.4205C19.3815 18.2095 19.5 17.9234 19.5 17.625V1.875C19.5 1.57663 19.3815 1.29048 19.1705 1.0795C18.9595 0.868526 18.6734 0.75 18.375 0.75ZM9.75 3H14.2289C14.6325 3 14.9789 3.31031 14.9991 3.71391C15.0039 3.81531 14.9882 3.91666 14.9528 4.0118C14.9173 4.10693 14.8629 4.19389 14.7929 4.26739C14.7229 4.34089 14.6386 4.3994 14.5453 4.43938C14.452 4.47937 14.3515 4.49999 14.25 4.5H9.77109C9.3675 4.5 9.02109 4.18969 9.00094 3.78609C8.99605 3.68469 9.01181 3.58334 9.04724 3.4882C9.08268 3.39307 9.13707 3.30611 9.2071 3.23261C9.27714 3.15911 9.36137 3.1006 9.45469 3.06062C9.54801 3.02063 9.64848 3.00001 9.75 3ZM8.41266 16.4916C8.10394 16.5252 7.79241 16.4621 7.52106 16.3111C7.2497 16.1601 7.03189 15.9287 6.89768 15.6486C6.76348 15.3686 6.71949 15.0538 6.7718 14.7477C6.82411 14.4416 6.97013 14.1593 7.18971 13.9397C7.40929 13.7201 7.69161 13.5741 7.9977 13.5218C8.3038 13.4695 8.61859 13.5135 8.89863 13.6477C9.17867 13.7819 9.41015 13.9997 9.56113 14.2711C9.71212 14.5424 9.77516 14.8539 9.74156 15.1627C9.70459 15.5024 9.55269 15.8194 9.31103 16.061C9.06936 16.3027 8.75242 16.4546 8.41266 16.4916ZM15.9127 16.4916C15.6039 16.5252 15.2924 16.4621 15.0211 16.3111C14.7497 16.1601 14.5319 15.9287 14.3977 15.6486C14.2635 15.3686 14.2195 15.0538 14.2718 14.7477C14.3241 14.4416 14.4701 14.1593 14.6897 13.9397C14.9093 13.7201 15.1916 13.5741 15.4977 13.5218C15.8038 13.4695 16.1186 13.5135 16.3986 13.6477C16.6787 13.7819 16.9101 13.9997 17.0611 14.2711C17.2121 14.5424 17.2752 14.8539 17.2416 15.1627C17.2046 15.5024 17.0527 15.8194 16.811 16.061C16.5694 16.3027 16.2524 16.4546 15.9127 16.4916ZM18 6.75V9.75H6V6.75H18Z'
                                        fill='#1E4BD2'
                                    />
                                    <path
                                        d='M13.9689 19.5L15.4393 21H8.56089L10.0314 19.5H7.96886L4.19214 23.25H6.31089L7.06089 22.5H16.9393L17.6893 23.25H19.8128L16.0782 19.5H13.9689Z'
                                        fill='#1E4BD2'
                                    />
                                </svg>

                                <Typography.Title
                                    style={{
                                        fontSize: 14,
                                        fontWeight: 'bold',
                                        margin: 0,
                                        marginLeft: 10,
                                    }}
                                    level={5}
                                >
                                    Ближайшая станция метро
                                </Typography.Title>
                            </>
                        }
                    >
                        <svg
                            width='21'
                            height='14'
                            viewBox='0 0 21 14'
                            fill='none'
                            xmlns='http://www.w3.org/2000/svg'
                        >
                            <g clipPath='url(#clip0_174_5232)'>
                                <path
                                    d='M1.40166 12.0442L6.16068 0L10.1863 7.04074L14.1956 0L18.9709 12.0442H20.3726V13.8696H13.1688V12.0442H14.2445L13.2014 9.0454L10.1863 14L7.17116 9.0454L6.12809 12.0442H7.20376V13.8696H3.24249e-05V12.0442H1.40166Z'
                                    fill='#F2782D'
                                />
                            </g>
                            <defs>
                                <clipPath id='clip0_174_5232'>
                                    <rect
                                        width='20.3725'
                                        height='14'
                                        fill='white'
                                        transform='matrix(-1 0 0 1 20.3726 0)'
                                    />
                                </clipPath>
                            </defs>
                        </svg>

                        <Typography.Text style={{ marginLeft: 10 }} className='base-text'>
                            Октябрьская
                        </Typography.Text>
                    </AccordionItem>
                    <AccordionItem
                        id='work-time'
                        className='department__details__item'
                        title={
                            <>
                                <svg
                                    width='24'
                                    height='24'
                                    viewBox='0 0 24 24'
                                    fill='none'
                                    xmlns='http://www.w3.org/2000/svg'
                                >
                                    <path
                                        d='M14.4347 11.6241L10.125 7.24219V11.3428C10.125 11.4174 10.1546 11.4889 10.2074 11.5417C10.2601 11.5944 10.3317 11.6241 10.4062 11.6241H14.4347Z'
                                        fill='#1E4BD2'
                                    />
                                    <path
                                        d='M8.625 12.5625V6.75H2.8125C2.66332 6.75 2.52024 6.80926 2.41475 6.91475C2.30926 7.02024 2.25 7.16332 2.25 7.3125V22.6875C2.25 22.8367 2.30926 22.9798 2.41475 23.0852C2.52024 23.1907 2.66332 23.25 2.8125 23.25H14.4375C14.5867 23.25 14.7298 23.1907 14.8352 23.0852C14.9407 22.9798 15 22.8367 15 22.6875V13.125H9.1875C9.03832 13.125 8.89524 13.0657 8.78975 12.9602C8.68426 12.8548 8.625 12.7117 8.625 12.5625Z'
                                        fill='#1E4BD2'
                                    />
                                    <path
                                        d='M17.1562 5.62406H21.1847L16.875 1.24219V5.34281C16.875 5.4174 16.9046 5.48894 16.9574 5.54169C17.0101 5.59443 17.0817 5.62406 17.1562 5.62406Z'
                                        fill='#1E4BD2'
                                    />
                                    <path
                                        d='M15.9375 7.125C15.7883 7.125 15.6452 7.06574 15.5398 6.96025C15.4343 6.85476 15.375 6.71168 15.375 6.5625V0.75H8.0625C7.91332 0.75 7.77024 0.809263 7.66475 0.914752C7.55926 1.02024 7.5 1.16332 7.5 1.3125V5.25H9.47437C9.72638 5.25101 9.97571 5.3018 10.208 5.39946C10.4403 5.49712 10.6511 5.63971 10.8281 5.81906L15.945 11.0222C16.302 11.3844 16.5013 11.873 16.4995 12.3816V18.75H21.1875C21.3367 18.75 21.4798 18.6907 21.5852 18.5852C21.6907 18.4798 21.75 18.3367 21.75 18.1875V7.125H15.9375Z'
                                        fill='#1E4BD2'
                                    />
                                </svg>

                                <Typography.Title
                                    style={{
                                        fontSize: 14,
                                        fontWeight: 'bold',
                                        margin: 0,
                                        marginLeft: 10,
                                    }}
                                    level={5}
                                >
                                    Обслуживание
                                </Typography.Title>
                            </>
                        }
                    >
                        <Row>
                            {department.special.person ? (
                                <>
                                    <svg
                                        width='14'
                                        height='12'
                                        viewBox='0 0 14 12'
                                        fill='none'
                                        xmlns='http://www.w3.org/2000/svg'
                                    >
                                        <path
                                            d='M7 5.72052C7.75211 5.72052 8.47342 5.42174 9.00524 4.88992C9.53707 4.35809 9.83584 3.63678 9.83584 2.88467C9.83584 2.13256 9.53707 1.41125 9.00524 0.879428C8.47342 0.347604 7.75211 0.0488281 7 0.0488281C6.24789 0.0488281 5.52658 0.347604 4.99476 0.879428C4.46293 1.41125 4.16416 2.13256 4.16416 2.88467C4.16416 3.63678 4.46293 4.35809 4.99476 4.88992C5.52658 5.42174 6.24789 5.72052 7 5.72052ZM7 6.96645C3.22956 6.96645 0.8125 9.04714 0.8125 10.0602V11.9519H13.1875V10.0602C13.1875 8.83508 10.8993 6.96645 7 6.96645Z'
                                            fill='#1E4BD2'
                                        />
                                    </svg>
                                    <Typography.Text
                                        style={{ marginLeft: 10 }}
                                        className='base-text'
                                    >
                                        Физические лица
                                    </Typography.Text>
                                </>
                            ) : null}
                        </Row>
                        <Row>
                            {department.special.juridical ? (
                                <>
                                    <svg
                                        width='14'
                                        height='12'
                                        viewBox='0 0 14 12'
                                        fill='none'
                                        xmlns='http://www.w3.org/2000/svg'
                                    >
                                        <path
                                            d='M13.75 2.8875C13.75 2.79799 13.7144 2.71215 13.6511 2.64885C13.5879 2.58556 13.502 2.55 13.4125 2.55H10.6V0.975C10.6 0.915326 10.5763 0.858097 10.5341 0.815901C10.4919 0.773705 10.4347 0.75 10.375 0.75H3.625C3.56533 0.75 3.5081 0.773705 3.4659 0.815901C3.42371 0.858097 3.4 0.915326 3.4 0.975V2.55H0.5875C0.497989 2.55 0.412145 2.58556 0.348851 2.64885C0.285558 2.71215 0.25 2.79799 0.25 2.8875L0.348851 5.25H13.8489L13.75 2.8875ZM9.475 2.55H4.525V1.875H9.475V2.55Z'
                                            fill='#1E4BD2'
                                        />
                                        <path
                                            d='M9.25 6.9H4.75V6H0.25V11.5125C0.25 11.602 0.285558 11.6879 0.348851 11.7511C0.412145 11.8144 0.497989 11.85 0.5875 11.85H13.4125C13.502 11.85 13.5879 11.8144 13.6511 11.7511C13.7144 11.6879 13.75 11.602 13.75 11.5125V6H9.25V6.9Z'
                                            fill='#1E4BD2'
                                        />
                                    </svg>

                                    <Typography.Text
                                        style={{ marginLeft: 10 }}
                                        className='base-text'
                                    >
                                        Юридические лица
                                    </Typography.Text>
                                </>
                            ) : null}
                        </Row>
                        <Row>
                            {department.special.vipOffice ? (
                                <>
                                    <svg
                                        width='16'
                                        height='8'
                                        viewBox='0 0 16 8'
                                        fill='none'
                                        xmlns='http://www.w3.org/2000/svg'
                                    >
                                        <path
                                            d='M12.8333 6.5H3.16667C2.95887 6.5 2.75662 6.47622 2.5625 6.4313V7.25C2.5625 7.56065 2.31066 7.8125 2 7.8125C1.68934 7.8125 1.4375 7.56065 1.4375 7.25V5.86348C0.86381 5.37433 0.5 4.6463 0.5 3.8333V2C0.5 1.17155 1.17157 0.5 2 0.5C2.82843 0.5 3.5 1.17155 3.5 2V2.9C3.5 3.23135 3.76863 3.5 4.1 3.5H11.9C12.2313 3.5 12.5 3.23135 12.5 2.9V2C12.5 1.17155 13.1715 0.5 14 0.5C14.8284 0.5 15.5 1.17155 15.5 2V3.8333C15.5 4.6463 15.1362 5.37433 14.5625 5.86348V7.25C14.5625 7.56065 14.3106 7.8125 14 7.8125C13.6893 7.8125 13.4375 7.56065 13.4375 7.25V6.4313C13.2434 6.47622 13.0411 6.5 12.8333 6.5Z'
                                            fill='#900158'
                                        />
                                    </svg>

                                    <Typography.Text
                                        style={{ marginLeft: 10 }}
                                        className='base-text'
                                    >
                                        Клиенты с пакетом "Привелегия"
                                    </Typography.Text>
                                </>
                            ) : null}
                        </Row>
                        <Row>
                            {department.special.Prime ? (
                                <>
                                    <svg
                                        width='16'
                                        height='8'
                                        viewBox='0 0 16 8'
                                        fill='none'
                                        xmlns='http://www.w3.org/2000/svg'
                                    >
                                        <path
                                            d='M12.8333 6.5H3.16667C2.95887 6.5 2.75662 6.47622 2.5625 6.4313V7.25C2.5625 7.56065 2.31066 7.8125 2 7.8125C1.68934 7.8125 1.4375 7.56065 1.4375 7.25V5.86348C0.86381 5.37433 0.5 4.6463 0.5 3.8333V2C0.5 1.17155 1.17157 0.5 2 0.5C2.82843 0.5 3.5 1.17155 3.5 2V2.9C3.5 3.23135 3.76863 3.5 4.1 3.5H11.9C12.2313 3.5 12.5 3.23135 12.5 2.9V2C12.5 1.17155 13.1715 0.5 14 0.5C14.8284 0.5 15.5 1.17155 15.5 2V3.8333C15.5 4.6463 15.1362 5.37433 14.5625 5.86348V7.25C14.5625 7.56065 14.3106 7.8125 14 7.8125C13.6893 7.8125 13.4375 7.56065 13.4375 7.25V6.4313C13.2434 6.47622 13.0411 6.5 12.8333 6.5Z'
                                            fill='#DF7139'
                                        />
                                    </svg>

                                    <Typography.Text
                                        style={{ marginLeft: 10 }}
                                        className='base-text'
                                    >
                                        Клиенты с пакетом "Прайм"
                                    </Typography.Text>
                                </>
                            ) : null}
                        </Row>
                    </AccordionItem>
                </Accordion>

                <Row>
                    <Typography.Title level={5}>Выберите время посещения офиса</Typography.Title>
                </Row>

                <Row style={{ gap: 10 }}>
                    {rootStore.selectedDepartmentDetails
                        ? rootStore.selectedDepartmentDetails.workload[0].loadHours.map(
                              (hour, index) => (
                                  <Button
                                      onClick={() => {
                                          setSelectedTimeIndex(index);
                                          setSelectedTime(hour.hour);
                                      }}
                                      style={{
                                          padding: 5,
                                          color: selectedTimeIndex === index ? '#1E4BD2' : '#333',
                                          background:
                                              selectedTimeIndex === index ? '#F3F7FA' : 'none',
                                          border: '1px solid #6B7683',
                                      }}
                                  >
                                      {hour.hour}
                                  </Button>
                              )
                          )
                        : null}
                </Row>
            </div>

            <div className='department__details__actions'>
                <Col>
                    <AdmiralButton
                        loading={isLoading}
                        onClick={() => {
                            setIsLoading(true);

                            rootStore
                                .createAppointment(selectedTime)
                                .then((ticket) => {
                                    api.info({
                                        message: 'Вы записаны на ' + selectedTime,
                                        description: (
                                            <div>
                                                <div>
                                                    Время на дорогу:{' '}
                                                    {Math.ceil(
                                                        ticket ? ticket?.estimatedTimeWalk / 60 : 0
                                                    )}{' '}
                                                    минут пешком. Или{' '}
                                                    {Math.ceil(
                                                        ticket ? ticket?.estimatedTimeCar / 60 : 0
                                                    )}{' '}
                                                    минут на машине
                                                </div>
                                            </div>
                                        ),
                                        placement: 'topRight',
                                    });
                                })
                                .catch((error) => {
                                    console.log(error);
                                    api.error({
                                        message: 'Ошибка записи',
                                        description: (
                                            <div>Попробуйте записаться на другое время</div>
                                        ),
                                        placement: 'topRight',
                                    });
                                })
                                .finally(() => {
                                    setIsLoading(false);
                                });
                        }}
                    >
                        Записаться в отделение
                    </AdmiralButton>
                    <a
                        href={`https://yandex.ru/maps/?ll=${rootStore.start[0]}%2C${rootStore.start[1]}&mode=routes&rtext=${rootStore.start[1]}%2C${rootStore.start[0]}~${rootStore.selectedDepartment?.location.coordinates.latitude}%2C${rootStore.selectedDepartment?.location.coordinates.longitude}&rtt=mt&ruri=~&z=14`}
                        target='_blank'
                        rel='noopener noreferrer'
                    >
                        <AdmiralButton appearance='secondary'>
                            Проложить маршрут в Яндекс Картах
                        </AdmiralButton>
                    </a>
                </Col>
            </div>
        </>
    );
});

export default DepartmentDetails;
