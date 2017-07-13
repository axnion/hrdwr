
const createDataMock = function() {
    let data = {}
    let cpus = []
    let disks = []
    let memory = {}
    let sensors = {}

    cpus.push({Name: "cpu0", Usage: createCPUFreq()})
    cpus.push({Name: "cpu1", Usage: createCPUFreq()})

    disks.push({Name: "/dev/sdc1", Total: 223390420, Used: createInt(20000, 223390400)})

    memory.push({Total: 12298264, Used: createInt(1000, 12298264)})

    let temps = []
    temps.push({Lable: "CPU Temperature", Value: createInt(20, 99)})
    temps.push({Lable: "MB Temperature", Value: createInt(20, 99)})
    temps.push({Lable: "Core 0", Value: createInt(20, 99)})
    temps.push({Lable: "Core 1", Value: createInt(20, 99)})

    let fans = []
    fans.push({Lable: "CPU FAN Speed", Value: createInt(0, 2500)})
    fans.push({Lable: "CHASSI1 FAN Speed", Value: createInt(0, 2500)})
    fans.push({Lable: "POWER FAN Speed", Value: createInt(0, 2500)})

    let volts = []
    volts.push({Lable: "Vcore Voltage", Value: createVolt()})
    volts.push({Lable: " +3.3 Voltage", Value: createVolt()})
    volts.push({Lable: " +5 Voltage", Value: createVolt()})
    volts.push({Lable: " +12 Voltage", Value: createVolt()})

    sensors.Temps = temps
    sensors.Fans = fans
    sensors.Volt = volts

}

const createCPUFreq = function() {
    return (Math.random() * (0.0001 - 1) + 1)
}

const createVolt = function() {
    return (Math.random() * (0.0001 - 12) + 12).toFixed(2)
}

const createInt = function(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

export default createDataMock